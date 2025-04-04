import { http, HttpResponse } from "msw";
import { API_URL } from "@/config/environment";

import { createUser } from "@tests/support/fixtures";

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
  
  http.post(`${API_URL}/api/v1/users`, ({ request }) => {
    const authHeader = request.headers.get("Authorization");
    if (!authHeader || !authHeader.startsWith("Bearer")) {
      return HttpResponse.json({ error: "Unauthorized" }, { status: 401 });
    }

    return HttpResponse.json({
      data: createUser(),
    });
  }),
];
