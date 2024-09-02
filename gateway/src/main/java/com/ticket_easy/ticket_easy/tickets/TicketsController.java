package com.ticket_easy.ticket_easy.tickets;

import com.ticket_easy.ticket_easy.infra.StandardError;
import com.ticket_easy.ticket_easy.tickets.dto.CreateTicketDTO;
import com.ticket_easy.ticket_easy.tickets.dto.CreateTicketToSendMicrosserviceDTO;
import com.ticket_easy.ticket_easy.tickets.dto.ResponseTicketDTO;
import com.ticket_easy.ticket_easy.users.domain.User;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.enums.SecuritySchemeType;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.security.SecurityScheme;
import io.swagger.v3.oas.annotations.security.SecuritySchemes;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import java.net.URI;

@RestController
@RequestMapping(value = "/tickets")
@ApiResponses(value = {
                @ApiResponse(responseCode = "400", description = "Erro de requisição", content = @Content(schema = @Schema(implementation = StandardError.class))),
                @ApiResponse(responseCode = "401", description = "Usuário não autenticado", content = @Content(schema = @Schema(implementation = StandardError.class))),
                @ApiResponse(responseCode = "403", description = "Não é possível acessar essa rota", content = @Content(schema = @Schema(implementation = StandardError.class))),
                @ApiResponse(responseCode = "500", description = "Erro desconhecido", content = @Content(schema = @Schema(implementation = StandardError.class))),
})
@SecuritySchemes({
                @SecurityScheme(name = "Authorization", type = SecuritySchemeType.HTTP, scheme = "bearer", bearerFormat = "JWT")
})
public class TicketsController {
        @Autowired
        private TicketsService ticketsService;

        // TODO: documentar melhor essa parte
        @Operation(summary = "Inicia o processo de garantir o ticket")
        @ApiResponses(value = {
                        @ApiResponse(responseCode = "201", description = "Informa que esta sendo feito o processo de garantir o ingresso, esse processo pode demorar um pouco pois ele é sincrono"),
        })
        @PostMapping
        public ResponseEntity<ResponseTicketDTO> create(@RequestBody CreateTicketDTO dto) {
                User authenticatedUser = (User) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
                CreateTicketToSendMicrosserviceDTO dtoSendMicroservice = new CreateTicketToSendMicrosserviceDTO();
                dtoSendMicroservice.setEventId(dto.getEventId());
                dtoSendMicroservice.setUserId(authenticatedUser.getId());

                ResponseTicketDTO responseTicketDTO = ticketsService.create(dtoSendMicroservice);
                URI uri = ServletUriComponentsBuilder
                        .fromCurrentRequest()
                        .path("/{id}")
                        .buildAndExpand(responseTicketDTO.getData().getId())
                        .toUri();

                return ResponseEntity
                        .created(uri)
                        .body(responseTicketDTO);
        }

        @Operation(summary = "Busca por um ticket pelo id")
        @ApiResponses(value = {
                        @ApiResponse(responseCode = "200", description = "Retorna o ingresso buscado"),
        })
        @GetMapping("/{id}")
        public ResponseEntity<ResponseTicketDTO> findById(@PathVariable String id) {
                User authenticatedUser = (User) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
                ResponseTicketDTO responseTicketDTO = ticketsService.findById(id, authenticatedUser.getId());
                return ResponseEntity.ok(responseTicketDTO);
        }
}
