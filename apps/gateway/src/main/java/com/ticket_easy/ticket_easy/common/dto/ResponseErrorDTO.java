package com.ticket_easy.ticket_easy.common.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
public class ResponseErrorDTO {
    private String message;
    private Long status;
}
