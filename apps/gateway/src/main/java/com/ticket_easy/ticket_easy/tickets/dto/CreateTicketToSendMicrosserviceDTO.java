package com.ticket_easy.ticket_easy.tickets.dto;

import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
public class CreateTicketToSendMicrosserviceDTO {
    private String userId;
    private String eventId;

}
