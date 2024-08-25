package com.ticket_easy.ticket_easy;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

@SpringBootApplication
@EnableJpaAuditing
public class TicketEasyApplication {

	public static void main(String[] args) {
		SpringApplication.run(TicketEasyApplication.class, args);
	}

}
