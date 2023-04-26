use rocket::http::Status;
use rocket::response::status;
use rocket::serde::json::Json;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct ApiKeyResponse {
    api_key: String,
}

#[post("/")]
pub fn generate_key() -> status::Custom<Json<ApiKeyResponse>> {
    let api_key = "12345678".to_string();
    let response = ApiKeyResponse { api_key };
    status::Custom(
        Status::Created,
        Json(response),
    )

}
