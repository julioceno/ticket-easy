package com.ticket_easy.ticket_easy.infra.auth;

import com.ticket_easy.ticket_easy.auth.dto.SubjectDTO;
import com.ticket_easy.ticket_easy.auth.services.ValidateTokenService;
import com.ticket_easy.ticket_easy.users.domain.User;
import com.ticket_easy.ticket_easy.users.domain.UserRepository;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.Collections;

@Component
public class SecurityFilter extends OncePerRequestFilter {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private ValidateTokenService validateTokenService;

    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        String token = recoverToken(request);
        SubjectDTO subject = validateToken(token);

        if (subject != null) {
            User user = userRepository.findByEmail(subject.getEmail()).orElse(null);
            var authorities = Collections.singletonList(new SimpleGrantedAuthority("ROLE_USER"));
            var authentication = new UsernamePasswordAuthenticationToken(user, null, authorities);
            SecurityContextHolder.getContext().setAuthentication(authentication);
        }

        filterChain.doFilter(request, response);
    }

    private String recoverToken(HttpServletRequest request) {
        String authorization = request.getHeader("Authorization");
        if (authorization == null) return null;
        return authorization.replace("Bearer ", "");
    }

    private SubjectDTO validateToken(String token) {
        try {
            return validateTokenService.run(token);
        } catch (RuntimeException exception) {
            return null;
        }
    }
}
