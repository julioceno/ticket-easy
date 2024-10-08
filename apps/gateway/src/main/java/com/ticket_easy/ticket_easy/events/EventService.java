package com.ticket_easy.ticket_easy.events;

import com.ticket_easy.ticket_easy.common.validations.IntegrationErrorMap;
import com.ticket_easy.ticket_easy.events.dto.*;
import com.ticket_easy.ticket_easy.exceptions.InternalServerErrorException;
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
public class EventService {
    private static final Logger logger = LoggerFactory.getLogger(EventService.class.getName());

    @Value("${api.integrations.events.url}")
    private String eventUrl;

    @Value("${api.integrations.events.secret}")
    private String secret;

    public ResponseEventsListDTO fetchEvents(QueryEventsDTO query) {
        logger.info("FETCH EVENTS: Creating request in event api...");
        HttpEntity<String> httpEntity = createHttpEntity();
        RestTemplate restTemplate = new RestTemplate();

        try {
            UriComponentsBuilder urlBuilder = UriComponentsBuilder.fromHttpUrl(getUrl());
            if (query.getName() != null) {
                urlBuilder.queryParam("name", query.getName());
            }

            if (query.getSkip() != null) {
                urlBuilder.queryParam("skip", query.getSkip());
            }

            if (query.getLimit() != null) {
                urlBuilder.queryParam("limit", query.getLimit());
            }

            ResponseEntity<ResponseEventsListDTO> response = restTemplate.exchange(
                    urlBuilder.toUriString(),
                    HttpMethod.GET,
                    httpEntity,
                    ResponseEventsListDTO.class,
                    query
            );

            logger.info("FETCH EVENTS: Response obtained");
            return response.getBody();
        } catch (HttpClientErrorException err) {
            IntegrationErrorMap.validation(err);
            throw new InternalServerErrorException("Ocorreu um erro deconhecido");
        }
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
            IntegrationErrorMap.validation(err);
            throw new InternalServerErrorException("Ocorreu um erro deconhecido");
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
