defmodule PulapWeb.AuditHelpers do
  @moduledoc """
  Helpers to set created_by and updated_by fields in entities and relations.
  """

  def maybe_put_created_by(params, conn) do
    case conn.assigns[:current_user] do
      %{} = user -> Map.put(params, "created_by", user.id)
      _ -> params
    end
  end

  def maybe_put_updated_by(params, conn) do
    case conn.assigns[:current_user] do
      %{} = user -> Map.put(params, "updated_by", user.id)
      _ -> params
    end
  end
end
