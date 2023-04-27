#[macro_use]
extern crate rocket;
use dotenv::dotenv;
mod controllers;
mod services;

#[launch]
fn rocket() -> _ {
    dotenv().ok();
    rocket::build().mount(
        "/api-key",
        routes![
            controllers::api_key_controller::generate_key,
            controllers::api_key_controller::validate_key
        ],
    )
}
