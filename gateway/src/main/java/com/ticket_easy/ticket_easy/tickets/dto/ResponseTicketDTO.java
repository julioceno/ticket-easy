package com.ticket_easy.ticket_easy.tickets.dto;


import com.fasterxml.jackson.annotation.JsonAlias;
import lombok.*;

import java.util.Date;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
public class TicketDTO {
    @JsonAlias("_id")
    private String id;
    private String userId;
    private String messageError;
    private String ticketPrice;
    private String dayEvent;
    private String eventName;
    private Date createdAt;
    private Date updatedAt;
}
