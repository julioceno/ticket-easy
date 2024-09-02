package com.ticket_easy.ticket_easy.events.dto;

import com.ticket_easy.ticket_easy.common.dto.QueryRequestDTO;
import lombok.*;

@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
public class QueryEventsDTO extends QueryRequestDTO {
    private String name;
}
