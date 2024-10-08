package com.ticket_easy.ticket_easy.infra;

import com.ticket_easy.ticket_easy.exceptions.BadRequestException;
import com.ticket_easy.ticket_easy.exceptions.InternalServerErrorException;
import com.ticket_easy.ticket_easy.exceptions.NotFoundException;
import com.ticket_easy.ticket_easy.exceptions.UnauthorizedException;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

import java.time.Instant;

@ControllerAdvice
public class ResourceExceptionHandler {
    @ExceptionHandler(BadRequestException.class)
    public ResponseEntity<StandardError> badRequest(
            BadRequestException e,
            HttpServletRequest request
    ) {
        String error = "Bad Request";
        HttpStatus status = HttpStatus.BAD_REQUEST;
        StandardError err = new StandardError(Instant.now(), status.value(), error, e.getMessage(), request.getRequestURI());
        return ResponseEntity.status(status).body(err);
    }

    @ExceptionHandler(NotFoundException.class)
    public ResponseEntity<StandardError> badRequest(
            NotFoundException e,
            HttpServletRequest request
    ) {
        String error = "Not Found";
        HttpStatus status = HttpStatus.NOT_FOUND;
        StandardError err = new StandardError(Instant.now(), status.value(), error, e.getMessage(), request.getRequestURI());
        return ResponseEntity.status(status).body(err);
    }

    @ExceptionHandler(UnauthorizedException.class)
    public ResponseEntity<StandardError> badRequest(
            UnauthorizedException e,
            HttpServletRequest request
    ) {
        String error = "Unauthorized";
        HttpStatus status = HttpStatus.UNAUTHORIZED;
        StandardError err = new StandardError(Instant.now(), status.value(), error, e.getMessage(), request.getRequestURI());
        return ResponseEntity.status(status).body(err);
    }

    @ExceptionHandler(InternalServerErrorException.class)
    public ResponseEntity<StandardError> badRequest(
            InternalServerErrorException e,
            HttpServletRequest request
    ) {
        String error = "Unknown Error";
        HttpStatus status = HttpStatus.INTERNAL_SERVER_ERROR;
        StandardError err = new StandardError(Instant.now(), status.value(), error, e.getMessage(), request.getRequestURI());
        return ResponseEntity.status(status).body(err);
    }
}
