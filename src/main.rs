#[macro_use]
extern crate rocket;
mod controllers;

#[launch]
fn rocket() -> _ {
    rocket::build().mount(
        "/api-key",
        routes![controllers::api_key_controller::generate_key],
    )
}
