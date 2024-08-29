package com.ticket_easy.ticket_easy.common.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
public class ResponseDTO<T> {
    @JsonProperty("message")
    private String message;

    @JsonProperty("status")
    private Long status;

    @JsonProperty("data")
    private T data;
}
