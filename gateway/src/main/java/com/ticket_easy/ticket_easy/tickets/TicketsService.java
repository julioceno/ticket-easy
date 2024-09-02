package com.ticket_easy.ticket_easy.tickets;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.ticket_easy.ticket_easy.common.validations.IntegrationErrorMap;
import com.ticket_easy.ticket_easy.exceptions.BadRequestException;
import com.ticket_easy.ticket_easy.exceptions.InternalServerErrorException;
import com.ticket_easy.ticket_easy.tickets.dto.CreateTicketDTO;
import com.ticket_easy.ticket_easy.tickets.dto.CreateTicketToSendMicrosserviceDTO;
import com.ticket_easy.ticket_easy.tickets.dto.ResponseTicketDTO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.util.UriComponentsBuilder;

@Service
public class TicketsService {
    private static final Logger logger = LoggerFactory.getLogger(TicketsService.class.getName());

    @Value("${api.integrations.tickets.url}")
    private String ticketsUrl;

    @Value("${api.integrations.tickets.secret}")
    private String secret;

    public ResponseTicketDTO create(CreateTicketToSendMicrosserviceDTO dto) {
        try {
            HttpEntity<String> httpEntity = createHttpEntity(dto);
            RestTemplate restTemplate = new RestTemplate();

            return restTemplate.postForEntity(
                    getUrl(),
                    httpEntity,
                    ResponseTicketDTO.class
            ).getBody();
        } catch (HttpClientErrorException err) {
            IntegrationErrorMap.validation(err);
            throw new InternalServerErrorException();
        }
    }

    public ResponseTicketDTO findById(String id, String userId) {
        try {
            HttpEntity<String> httpEntity = createHttpEntity(null);
            RestTemplate restTemplate = new RestTemplate();

            UriComponentsBuilder urlBuilder = UriComponentsBuilder.fromHttpUrl(getUrl() + id);
            urlBuilder.queryParam("userId", userId);

            ResponseEntity<ResponseTicketDTO> response = restTemplate.exchange(
                    urlBuilder.toUriString(),
                    HttpMethod.GET,
                    httpEntity,
                    ResponseTicketDTO.class);

            return response.getBody();
        } catch (HttpClientErrorException err) {
            IntegrationErrorMap.validation(err);
            throw new InternalServerErrorException();
        }
    }

    private String getUrl() {
        return ticketsUrl + "/tickets/";
    }

    private HttpEntity<String> createHttpEntity(Object dto) {
        try {
            HttpHeaders headers = new HttpHeaders();
            headers.set("x-api-key", secret);

            ObjectMapper objectMapper = new ObjectMapper();
            String json = objectMapper.writeValueAsString(dto);

            return new HttpEntity<>(json, headers);
        } catch (JsonProcessingException err) {
            throw new BadRequestException("Ocorre um erro ao tentar garantir o ingresso");
        }

    }
}
