package com.ticket_easy.ticket_easy.common.validations;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.ticket_easy.ticket_easy.common.dto.ResponseErrorDTO;
import com.ticket_easy.ticket_easy.exceptions.BadRequestException;
import com.ticket_easy.ticket_easy.exceptions.InternalServerErrorException;
import com.ticket_easy.ticket_easy.exceptions.NotFoundException;
import com.ticket_easy.ticket_easy.exceptions.UnauthorizedException;
import org.springframework.http.HttpStatus;
import org.springframework.web.client.HttpClientErrorException;


public class IntegrationErrorMap {
    public static void validation(HttpClientErrorException  response) {
        try {
            int status = response.getStatusCode().value();
            String err = response.getResponseBodyAsString();
            ResponseErrorDTO responseErrorDTO = new ObjectMapper().readValue(err, ResponseErrorDTO.class);

            if (status == HttpStatus.NOT_FOUND.value()) {
                throw new NotFoundException(responseErrorDTO.getMessage());
            }

            if (status == HttpStatus.BAD_REQUEST.value()) {
                throw new BadRequestException(responseErrorDTO.getMessage());
            }

            if (status == HttpStatus.UNAUTHORIZED.value()) {
                throw new UnauthorizedException(responseErrorDTO.getMessage());
            }
        } catch (JsonProcessingException err) {
            throw new InternalServerErrorException("Ocorreu um erro deconhecido");
        }
    }
}
