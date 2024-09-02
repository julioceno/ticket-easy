package com.ticket_easy.ticket_easy.auth.services;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTVerificationException;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.ticket_easy.ticket_easy.auth.dto.SubjectDTO;
import com.ticket_easy.ticket_easy.exceptions.UnauthorizedException;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

@Service
public class ValidateTokenService {
    @Value("${api.security.token.secret}")
    private String secret;

    public SubjectDTO run(String token) {
        try {
            Algorithm algorithm = Algorithm.HMAC256(secret);
            JWTVerifier verifier = JWT.require(algorithm)
                    .withIssuer("ticket-easy")
                    .build();

            String subject = verifier.verify(token).getSubject();
            return new ObjectMapper().readValue(subject, SubjectDTO.class);
        } catch (JWTVerificationException | JsonProcessingException exception) {
            throw new UnauthorizedException();
        }
    }
}
