#[macro_use]
extern crate rocket;
mod controllers;

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/get-token", routes![controllers::token_controller::get_token])
}
