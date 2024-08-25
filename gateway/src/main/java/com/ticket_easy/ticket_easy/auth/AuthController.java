package com.ticket_easy.ticket_easy.auth;

import com.ticket_easy.ticket_easy.auth.dto.SignInDTO;
import com.ticket_easy.ticket_easy.auth.dto.TokensDTO;
import com.ticket_easy.ticket_easy.auth.services.AuthService;
import com.ticket_easy.ticket_easy.infra.StandardError;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/auth")
public class AuthController {
    @Autowired
    private AuthService authService;

    @PostMapping("sign-in")
    @Operation(summary = "Faz login do usuário")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Retorna os tokens de autorização do usuário"),
            @ApiResponse(responseCode = "401", description = "Usuário não pode ser autenticado", content = @Content(schema = @Schema(implementation = StandardError.class))),
    })
    public ResponseEntity<TokensDTO> signIn(@RequestBody SignInDTO dto) {
        TokensDTO tokensDTO = authService.signIn(dto);
        return ResponseEntity.ok(tokensDTO);
    }
}
