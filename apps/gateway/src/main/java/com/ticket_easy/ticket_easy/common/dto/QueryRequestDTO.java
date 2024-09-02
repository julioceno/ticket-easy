package com.ticket_easy.ticket_easy.common.dto;

import lombok.*;

@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
public class QueryRequestDTO {
    private Integer skip;
    private Integer limit;
}
