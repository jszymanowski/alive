import { createUser } from "@tests/support/fixtures";
import { HttpResponse, http } from "msw";
import { CreateUserRequestBody } from "@/api/userApi";
import { API_URL } from "@/config/environment";

export const handlers = [
  http.get(`${API_URL}/api/v1/current_user`, ({ request }) => {
    const authHeader = request.headers.get("Authorization");
    if (!authHeader || !authHeader.startsWith("Bearer")) {
      return HttpResponse.json({ error: "Unauthorized" }, { status: 401 });
    }

    return HttpResponse.json({
      data: createUser(),
    });
  }),

  http.post(`${API_URL}/api/v1/users`, async ({ request }) => {
    const authHeader = request.headers.get("Authorization");
    const body = (await request.json()) as CreateUserRequestBody;

    if (!body.name || !body.email) {
      return HttpResponse.json({ error: "Bad Request" }, { status: 400 });
    }

    if (!authHeader || !authHeader.startsWith("Bearer")) {
      return HttpResponse.json({ error: "Unauthorized" }, { status: 401 });
    }

    return HttpResponse.json({
      data: createUser({ name: body.name, email: body.email }),
    });
  }),
];
