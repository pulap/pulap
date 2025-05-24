defmodule PulapWeb.HomeController do
  use PulapWeb, :controller

  plug :put_layout, {PulapWeb.Layouts, :app}

  def show(conn, _params) do
    render(conn, :show)
  end
end
