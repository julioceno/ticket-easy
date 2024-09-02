package com.ticket_easy.ticket_easy.users.services;

import com.ticket_easy.ticket_easy.exceptions.BadRequestException;
import com.ticket_easy.ticket_easy.exceptions.NotFoundException;
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
public class FindOneUserService {
    private static final Logger logger = LoggerFactory.getLogger(FindOneUserService.class.getName());

    @Autowired
    private UserRepository userRepository;

    public UserDTO run(String id) {
        logger.info(format("Searching user by id %s...", id));
        User user = userRepository.findById(id).orElse(null);

        if (user == null) {
            logger.error(format("User with id %s not exists", id));
            throw new NotFoundException("Usuário não existe");
        }

        return new UserDTO(user);
    }
}
