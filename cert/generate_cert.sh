openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./nginx.key -out ./nginx.cert  -subj '/CN=nginx'