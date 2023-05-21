use actix_web::{get, App, HttpRequest, HttpResponse, HttpServer, Responder};
mod utils;

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world!")
}

#[get("/token")]
async fn token(info: HttpRequest) -> impl Responder {
    if let Some(auth_header) = info.headers().get("Authorization") {
        let token = auth_header.to_str().unwrap().replace("Bearer ", "");
        let decoded_token = utils::decode_token::decode_token(&token);
        let email = decoded_token.unwrap().email;
        return HttpResponse::Ok().body(email);
    }

    HttpResponse::Unauthorized().finish()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(hello).service(token))
        .bind(("0.0.0.0", 8000))?
        .run()
        .await
}
