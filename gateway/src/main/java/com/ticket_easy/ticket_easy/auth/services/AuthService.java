package com.ticket_easy.ticket_easy.auth.services;

import com.ticket_easy.ticket_easy.auth.dto.SignInDTO;
import com.ticket_easy.ticket_easy.auth.dto.TokensDTO;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@AllArgsConstructor
public class AuthService {
    private SignInService signInService;

    public TokensDTO signIn(SignInDTO signInDTO) {
        return signInService.run(signInDTO);
    }
}
