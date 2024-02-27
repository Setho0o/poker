package handler

import (
  "net/http"
  "web_proj/view/home"
)

func HandlerHomeIndex(w http.ResponseWriter, r *http.Request) {
  home.Index().Render(r.Context(), w)
}
