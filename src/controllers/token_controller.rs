use rocket::get;

#[get("/")]
pub fn get_token() -> &'static str {
    "12345678"
}
