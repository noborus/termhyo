# Security Policy

## Supported Versions

We provide security updates for the following versions of termhyo:

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability in termhyo, please report it privately to help us address it quickly.

### How to Report

1. **Email**: Send details to the maintainer via GitHub
2. **GitHub Security Advisory**: Use GitHub's private vulnerability reporting feature
3. **Do not** open a public issue for security vulnerabilities

### What to Include

When reporting a vulnerability, please include:

- Description of the vulnerability
- Steps to reproduce the issue
- Potential impact
- Any suggested fixes or mitigations
- Your contact information for follow-up

### Response Process

1. **Acknowledgment**: We will acknowledge receipt within 48 hours
2. **Investigation**: We will investigate and assess the impact
3. **Fix Development**: We will develop and test a fix
4. **Disclosure**: We will coordinate disclosure with you
5. **Release**: We will release a patched version

### Timeline

- Initial response: Within 48 hours
- Status update: Within 1 week
- Fix timeline: Depends on severity and complexity

## Security Considerations

When using termhyo:

- **Input Validation**: Always validate user input before passing to termhyo
- **Output Sanitization**: Be cautious when displaying user-controlled content
- **ANSI Injection**: Be aware that ANSI escape sequences can be injected through user input
- **Resource Usage**: Large tables or many columns may consume significant memory

## Best Practices

- Validate and sanitize all user input
- Limit table size when processing untrusted data
- Consider disabling ANSI escape sequences for untrusted content
- Use appropriate output encoding for your environment

Thank you for helping keep termhyo secure!
