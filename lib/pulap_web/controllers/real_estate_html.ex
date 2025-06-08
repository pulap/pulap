defmodule PulapWeb.RealEstateHTML do
  use PulapWeb, :html

  embed_templates "real_estate_html/*"

  def slug(estate), do: Pulap.Estate.RealEstate.slug(estate)
end
