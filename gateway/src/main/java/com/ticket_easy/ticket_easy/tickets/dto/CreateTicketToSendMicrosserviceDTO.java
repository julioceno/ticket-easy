package com.ticket_easy.ticket_easy.tickets.dto;

import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode
// TODO: refactor name
public class CreateTicketToSendMicrosserviceDTO {
    private String userId;
    private String eventId;

}
