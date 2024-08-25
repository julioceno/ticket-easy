package com.ticket_easy.ticket_easy.auth.dto;

import com.ticket_easy.ticket_easy.users.domain.User;
import lombok.*;

public record SignInDTO(String email, String password){
}
