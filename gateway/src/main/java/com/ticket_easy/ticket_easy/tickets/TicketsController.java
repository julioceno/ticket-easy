package com.ticket_easy.ticket_easy.tickets;

import com.ticket_easy.ticket_easy.infra.StandardError;
import com.ticket_easy.ticket_easy.tickets.dto.RequestTicketDTO;
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
                        @ApiResponse(responseCode = "200", description = "Informa que esta sendo feito o processo de tentar garantir o ingresso, esse processo pode demorar um pouco pois ele é sincrono"),
        })
        @PostMapping
        public ResponseEntity<ResponseTicketDTO> create(RequestTicketDTO dto) {
                // TODO: fazer essa funcao retornar um 201
                User authenticatedUser = (User) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
                dto.setUserId(authenticatedUser.getId());

                ResponseTicketDTO responseTicketDTO = ticketsService.create(dto);
                return ResponseEntity.ok(responseTicketDTO);
        }

        @Operation(summary = "Busca por um ticket pelo id")
        @ApiResponses(value = {
                        @ApiResponse(responseCode = "200", description = "Retorna o ingresso buscado"),
        })
        @GetMapping("/{id}")
        public ResponseEntity<ResponseTicketDTO> findById(@PathVariable String id) {
                ResponseTicketDTO responseTicketDTO = ticketsService.findById(id);
                return ResponseEntity.ok(responseTicketDTO);
        }
}
