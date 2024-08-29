package com.ticket_easy.ticket_easy.events;

import com.ticket_easy.ticket_easy.common.validations.IntgrationValidation;
import com.ticket_easy.ticket_easy.events.dto.*;
import com.ticket_easy.ticket_easy.users.services.FindOneUserService;
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

import static java.lang.String.format;

@Service
public class EventService {
    private static final Logger logger = LoggerFactory.getLogger(EventService.class.getName());

    @Value("${api.integrations.event.url}")
    private String eventUrl;

    @Value("${api.integrations.event.secret}")
    private String secret;

    public ResponseEventsListDTO fetchEvents() {
        logger.info("FETCH EVENTS: Creating request in event api...");
        HttpEntity<String> httpEntity = createHttpEntity();
        RestTemplate restTemplate = new RestTemplate();

        ResponseEntity<ResponseEventsListDTO> response = restTemplate.exchange(
                getUrl(),
                HttpMethod.GET,
                httpEntity,
                ResponseEventsListDTO.class
        );

        // TODO: handler errors

        logger.info("FETCH EVENTS: Response obtained");
        return response.getBody();
    }

    public ResponseEventDTO fetchEvent(String id) {
        logger.info("FETCH: Creating request in event api...");
        HttpEntity<String> httpEntity = createHttpEntity();
        RestTemplate restTemplate = new RestTemplate();

        try {
            String urlWithId = getUrl() + "/{id}";
            ResponseEntity<ResponseEventDTO> response = restTemplate.exchange(
                    urlWithId,
                    HttpMethod.GET,
                    httpEntity,
                    ResponseEventDTO.class,
                    id
            );

            logger.info("FETCH: Response obtained");
            return response.getBody();
        } catch (HttpClientErrorException err) {
            IntgrationValidation.validation(err);
            throw new RuntimeException();
        }
    }


    private String getUrl() {
        return eventUrl + "/events";
    }

    private HttpEntity<String> createHttpEntity() {
        HttpHeaders headers = new HttpHeaders();
        headers.set("x-api-key", secret);
        return new HttpEntity<>(headers);
    }
}
