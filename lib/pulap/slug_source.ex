defprotocol Pulap.SlugSource do
  @moduledoc """
  Protocol defining how to extract the source text for a slug from a struct.
  """

  @doc """
  Returns the string to be used as the base for the slug.
  """
  @spec source_for_slug(struct) :: String.t()
  def source_for_slug(struct)
end
