package com.ticket_easy.ticket_easy.auth.services;

import com.ticket_easy.ticket_easy.auth.dto.SignInDTO;
import com.ticket_easy.ticket_easy.auth.dto.TokensDTO;
import com.ticket_easy.ticket_easy.exceptions.UnauthorizedException;
import com.ticket_easy.ticket_easy.users.domain.User;
import com.ticket_easy.ticket_easy.users.domain.UserRepository;
import com.ticket_easy.ticket_easy.users.dto.UserDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCrypt;
import org.springframework.stereotype.Service;

import static java.lang.String.format;

@Service
public class SignInService {
    private static final Logger logger = LoggerFactory.getLogger(SignInService.class.getName());

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private CreateTokenService createTokenService;

    public TokensDTO run(SignInDTO signInDTO) {
        User user = getUserByEmail(signInDTO.email());
        this.validatePassword(user.getPassword(), signInDTO.password());

        String token = createToken(user);
        return new TokensDTO(token);
    }

    private User getUserByEmail(String email) {
        logger.info(format("Searching user by email %s...", email));
        User user = userRepository.findByEmail(email).orElse(null);

        if (user == null) {
            logger.error("User not exists");
            throw new UnauthorizedException();
        }

        return user;
    }

    private void validatePassword(String userPassword, String enteredPassword) {
        logger.info("Validating password...");
        boolean isValidPassword = BCrypt.checkpw(enteredPassword, userPassword);

        if (!isValidPassword) {
            logger.error("Password is invalid");
            throw new UnauthorizedException();
        }

        logger.info("Password are valid");
    }

    private String createToken(User user) {
        logger.info("Generating token...");
        String accessToken = createTokenService.run(user);
        logger.info("Token generated");

        return accessToken;
    }

}
