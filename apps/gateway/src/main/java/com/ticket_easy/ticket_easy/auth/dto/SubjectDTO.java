package com.ticket_easy.ticket_easy.auth.dto;

import com.ticket_easy.ticket_easy.users.domain.User;
import lombok.*;

@Getter
@Setter
@EqualsAndHashCode
@AllArgsConstructor
@NoArgsConstructor
public class SubjectDTO {
    private String id;
    private String email;

    public SubjectDTO(User user) {
        this.id = user.getId();
        this.email = user.getEmail();
    }
}
