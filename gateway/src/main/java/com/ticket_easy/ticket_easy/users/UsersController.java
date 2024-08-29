package com.ticket_easy.ticket_easy.users;

import com.ticket_easy.ticket_easy.infra.StandardError;
import com.ticket_easy.ticket_easy.users.domain.User;
import com.ticket_easy.ticket_easy.users.dto.CreateUserDTO;
import com.ticket_easy.ticket_easy.users.dto.UserDTO;
import com.ticket_easy.ticket_easy.users.services.UsersService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.enums.ParameterIn;
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
@RequestMapping(value = "/users")
@ApiResponses(value = {
        @ApiResponse(responseCode = "400",  description = "Erro de requisição", content = @Content(schema = @Schema(implementation = StandardError.class))),
        @ApiResponse(responseCode = "401", description = "Usuário não autenticado", content = @Content(schema = @Schema(implementation = StandardError.class))),
        @ApiResponse(responseCode = "403", description = "Não é possível acessar essa rota", content = @Content(schema = @Schema(implementation = StandardError.class))),
        @ApiResponse(responseCode = "500", description = "Erro desconhecido", content = @Content(schema = @Schema(implementation = StandardError.class))),
})
@SecuritySchemes({
        @SecurityScheme(name = "Authorization", type = SecuritySchemeType.HTTP, scheme = "bearer", bearerFormat = "JWT")
})
public class UsersController {
    @Autowired
    private UsersService userService;

    @PostMapping
    @Operation(summary = "Cria o usuário")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Criação do usuário feita com sucesso"),
            @ApiResponse(responseCode = "400",  description = "Email do usuário já existe ou erro de requisição", content = @Content(schema = @Schema(implementation = StandardError.class))),
    })
    public ResponseEntity<UserDTO> create(@RequestBody CreateUserDTO dto) {
        UserDTO userDTO = userService.create(dto);
        URI uri = ServletUriComponentsBuilder
                .fromCurrentRequest()
                .path("/{id}")
                .buildAndExpand(userDTO.getId())
                .toUri();

        return ResponseEntity
                .created(uri)
                .body(userDTO);
    }

    @GetMapping("/me")
    @Operation(summary = "Retorna os dados do próprio usuário que esta chamando")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Retorna os dados do usuário"),
    })
    public ResponseEntity<UserDTO> findMe() {
        User authenticatedUser = (User) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        UserDTO userDTO = userService.findById(authenticatedUser.getId());
        return ResponseEntity.ok(userDTO);
    }
}
