package com.ticket_easy.ticket_easy.events;


import com.ticket_easy.ticket_easy.events.dto.ResponseEventDTO;
import com.ticket_easy.ticket_easy.events.dto.ResponseEventsListDTO;
import com.ticket_easy.ticket_easy.infra.StandardError;
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
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(value = "/events")
@ApiResponses(value = {
        @ApiResponse(responseCode = "400",  description = "Erro de requisição", content = @Content(schema = @Schema(implementation = StandardError.class))),
        @ApiResponse(responseCode = "403", description = "Usuário não autenticado", content = @Content(schema = @Schema(implementation = StandardError.class))),
        @ApiResponse(responseCode = "500", description = "Erro desconhecido", content = @Content(schema = @Schema(implementation = StandardError.class))),
})
@SecuritySchemes({
        @SecurityScheme(name = "Authorization", type = SecuritySchemeType.HTTP, scheme = "bearer", bearerFormat = "JWT")
})
public class EventController {
    @Autowired
    private EventService eventService;

    @GetMapping
    @Operation(summary = "Busca por eventos")
    public ResponseEntity<ResponseEventsListDTO> fetchEvents() {
        ResponseEventsListDTO response = eventService.fetchEvents();
        return ResponseEntity.ok(response);
    }

    @GetMapping("/{id}")
    @Operation(summary = "Busca um evento pelo id")
    public ResponseEntity<ResponseEventDTO> fetchEvent(@PathVariable String id) {
        ResponseEventDTO response = eventService.fetchEvent(id);
        return ResponseEntity.ok(response);
    }
}
