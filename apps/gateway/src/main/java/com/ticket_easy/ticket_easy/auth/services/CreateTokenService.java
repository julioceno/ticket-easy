package com.ticket_easy.ticket_easy.auth.services;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTVerificationException;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.ticket_easy.ticket_easy.auth.dto.SubjectDTO;
import com.ticket_easy.ticket_easy.exceptions.UnauthorizedException;
import com.ticket_easy.ticket_easy.users.domain.User;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.time.Duration;
import java.time.Instant;

@Service
public class CreateTokenService {
    @Value("${api.security.token.secret}")
    private String secret;

    public String run(User user) {
        try {
            SubjectDTO subjectDTO = new SubjectDTO(user);
            ObjectMapper objectMapper = new ObjectMapper();
            String subjectJson = objectMapper.writeValueAsString(subjectDTO);

            Algorithm algorithm = Algorithm.HMAC256(secret);
            return JWT.create()
                    .withIssuer("ticket-easy")
                    .withExpiresAt(this.generateExpirationDate())
                    .withSubject(subjectJson)
                    .sign(algorithm);
        } catch (JWTVerificationException | JsonProcessingException exception) {
            throw new UnauthorizedException();
        }
    }

    private Instant generateExpirationDate() {
        return Instant
                .now()
                .plus(Duration.ofHours(30)); // TODO: add in settings
    }
}
