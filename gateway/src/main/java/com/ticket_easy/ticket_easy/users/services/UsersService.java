package com.ticket_easy.ticket_easy.users.services;

import com.ticket_easy.ticket_easy.users.dto.CreateUserDTO;
import com.ticket_easy.ticket_easy.users.dto.UserDTO;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@AllArgsConstructor
public class UsersService {
    private final CreateUserService createUserService;
    private final FindOneUserService findOneUserService;

    public UserDTO create(CreateUserDTO createUserDTO) {
        return createUserService.run(createUserDTO);
    }

    public UserDTO findById(String id) {
        return findOneUserService.run(id);
    }

}
