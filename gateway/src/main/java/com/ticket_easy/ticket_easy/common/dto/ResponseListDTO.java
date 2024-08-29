package com.ticket_easy.ticket_easy.common.dto;

import lombok.*;

import java.util.ArrayList;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
public class ResponseListDTO<T> {
    private Long count;
    private ArrayList<T> data;
}
