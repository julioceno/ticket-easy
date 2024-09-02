package com.ticket_easy.ticket_easy.common.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
public class ResponseDTO<T> {
    private String message;
    private Long status;
    private T data;
}
