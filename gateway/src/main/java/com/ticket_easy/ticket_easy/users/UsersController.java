package com.ticket_easy.ticket_easy.users;

import com.ticket_easy.ticket_easy.auth.dto.SubjectDTO;
import com.ticket_easy.ticket_easy.users.domain.User;
import com.ticket_easy.ticket_easy.users.dto.CreateUserDTO;
import com.ticket_easy.ticket_easy.users.dto.UserDTO;
import com.ticket_easy.ticket_easy.users.services.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import java.net.URI;

@RestController
@RequestMapping(value = "/users")
public class UsersController {
    @Autowired
    private UsersService userService;

    @PostMapping
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
    public ResponseEntity<UserDTO> findById() {
        User authenticatedUser = (User) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        UserDTO userDTO = userService.findById(authenticatedUser.getId());
        return ResponseEntity.ok(userDTO);
    }
}
