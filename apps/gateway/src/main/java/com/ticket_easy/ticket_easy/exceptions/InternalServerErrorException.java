package com.ticket_easy.ticket_easy.exceptions;

public class InternalServerErrorException extends RuntimeException {
    public InternalServerErrorException(String msg) {
        super(msg);
    }

    public InternalServerErrorException() {
        super("Ocorreu um erro deconhecido");
    }
}