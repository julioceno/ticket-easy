package com.ticket_easy.ticket_easy.users.services;

import com.ticket_easy.ticket_easy.exceptions.BadRequestException;
import com.ticket_easy.ticket_easy.users.domain.User;
import com.ticket_easy.ticket_easy.users.domain.UserRepository;
import com.ticket_easy.ticket_easy.users.dto.CreateUserDTO;
import com.ticket_easy.ticket_easy.users.dto.UserDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import static java.lang.String.format;

@Service
public class CreateUserService {
    private static final Logger logger = LoggerFactory.getLogger(CreateUserService.class.getName());

    @Autowired
    private UserRepository userRepository;

    public UserDTO run(CreateUserDTO dto) {
        verifyIfUserAlreadyExists(dto.email());
        User userCreated = createUser(dto);

        return new UserDTO(userCreated);
    }

    private void verifyIfUserAlreadyExists(String email) {
        logger.info(format("Search user by email %s...", email));
        User user = userRepository.findByEmail(email).orElse(null);

        if (user != null) {
            logger.error("User already exists");
            throw new BadRequestException("Usuário já existente.");
        }

        logger.info(format("User with email %s not exists", email));
    }

    private User createUser(CreateUserDTO dto) {
        User userBuilded = buildUser(dto);

        logger.info("Creating user...");
        User userCreated = userRepository.save(userBuilded);
        logger.info("User created");

        return userCreated;
    }

    private User buildUser(CreateUserDTO dto) {
        logger.info("Building user...");
        String encyptedPassword = new BCryptPasswordEncoder().encode(dto.password());

        User newUser = new User();
        newUser.setName(dto.name());
        newUser.setEmail(dto.email());
        newUser.setPassword(encyptedPassword);

        logger.info("User builded");
        return newUser;
    }

}
