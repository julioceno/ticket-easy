package com.ticket_easy.ticket_easy.common.validations;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.ticket_easy.ticket_easy.auth.dto.SubjectDTO;
import com.ticket_easy.ticket_easy.exceptions.NotFoundException;
import org.springframework.http.HttpStatus;
import org.springframework.web.client.HttpClientErrorException;


public class IntgrationValidation {

    public static void validation(HttpClientErrorException response) throws JsonProcessingException {
        int status = response.getStatusCode().value();
        String err = response.getResponseBodyAsString();
         new ObjectMapper().readValue(err, SubjectDTO.class);

        if (status == HttpStatus.NOT_FOUND.value()) {
            throw new NotFoundException("");
        }
    }
}
