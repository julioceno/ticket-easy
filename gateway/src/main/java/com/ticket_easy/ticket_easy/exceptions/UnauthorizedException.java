package com.ticket_easy.ticket_easy.exceptions;

public class UnauthorizedException extends RuntimeException {
    public UnauthorizedException(String msg) {
        super(msg);
    }

    public UnauthorizedException() {
        super("Usuário não autorizado.");
    }
}