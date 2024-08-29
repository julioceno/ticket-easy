package com.ticket_easy.ticket_easy.events.dto;

import com.fasterxml.jackson.annotation.JsonAlias;
import lombok.*;

import java.util.Date;
import java.util.List;

@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
@EqualsAndHashCode(of = "id")
public class EventDTO {
    @JsonAlias("_id")
    private String id;
    private String name;
    private String description;
    private Double ticketValue;
    private List<String> imagesUrl;
    private Integer quantityTickets;
    private Date occuredAt;
}
