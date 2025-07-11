defmodule PulapWeb.RealEstateFormLive do
  use PulapWeb, :live_view

  on_mount {PulapWeb.UserAuth, :ensure_authenticated}

  import Ecto.Query
  alias Pulap.Repo
  alias Pulap.Set.Option
  alias Pulap.Estate.RealEstate

  def mount(_params, _session, socket) do
    categories = get_options_by_set("estate_category")

    socket =
      socket
      |> assign(live_action: :new)
      |> assign(:categories, categories)
      |> assign(:types, [])
      |> assign(:subtypes, [])
      |> assign(:selected_category, nil)
      |> assign(:selected_type, nil)
      |> assign(:selected_subtype, nil)
      |> assign(:changeset, RealEstate.changeset(%RealEstate{}, %{}))
      |> assign(:show_manual_coords, false)

    {:ok, socket}
  end

  def handle_event("address_selected", %{"lat" => lat, "lng" => lng, "address" => address_data}, socket) do
    
    changeset =
      socket.assigns.changeset
      |> Ecto.Changeset.put_change(:lat, String.to_float(lat))
      |> Ecto.Changeset.put_change(:lng, String.to_float(lng))
      |> Ecto.Changeset.put_change(:street, Map.get(address_data["address"], "road"))
      |> Ecto.Changeset.put_change(:number, Map.get(address_data["address"], "house_number"))
      |> Ecto.Changeset.put_change(:floor, Map.get(address_data["address"], "floor"))
      |> Ecto.Changeset.put_change(:apartment, Map.get(address_data["address"], "unit"))
      |> Ecto.Changeset.put_change(:postal_code, Map.get(address_data["address"], "postcode"))
      |> Ecto.Changeset.put_change(:admin_level_0, Map.get(address_data["address"], "country"))
      |> Ecto.Changeset.put_change(:admin_level_1, Map.get(address_data["address"], "state"))
      |> Ecto.Changeset.put_change(:admin_level_2, Map.get(address_data["address"], "county"))
      |> Ecto.Changeset.put_change(:admin_level_3, Map.get(address_data["address"], "city", Map.get(address_data["address"], "town", Map.get(address_data["address"], "village"))))
      |> Ecto.Changeset.put_change(:admin_level_4, Map.get(address_data["address"], "suburb"))

    
    {:noreply, assign(socket, changeset: changeset, show_manual_coords: true)}
  end

  def handle_event("toggle_geo_data_manual", _params, socket) do
    
    {:noreply, assign(socket, :show_manual_coords, not socket.assigns.show_manual_coords)}
  end

  def handle_event("select_category", %{"real_estate" => %{"category" => category_label}}, socket) do
    is_selected = category_label != ""
    types = if is_selected, do: get_options_by_parent_label("estate_type", category_label), else: []

    changeset =
      socket.assigns.changeset
      |> Ecto.Changeset.put_change(:category, category_label)
      |> Ecto.Changeset.put_change(:type, nil)
      |> Ecto.Changeset.put_change(:subtype, nil)

    {:noreply,
     socket
     |> assign(:selected_category, if(is_selected, do: category_label, else: nil))
     |> assign(:types, types)
     |> assign(:selected_type, nil)
     |> assign(:subtypes, [])
     |> assign(:selected_subtype, nil)
     |> assign(:changeset, changeset)}
  end

  def handle_event("select_type", %{"real_estate" => %{"type" => type_label}}, socket) do
    is_selected = type_label != ""
    subtypes = if is_selected, do: get_options_by_parent_label("estate_subtype", type_label), else: []

    changeset =
      socket.assigns.changeset
      |> Ecto.Changeset.put_change(:type, type_label)
      |> Ecto.Changeset.put_change(:subtype, nil)

    {:noreply,
     socket
     |> assign(:selected_type, if(is_selected, do: type_label, else: nil))
     |> assign(:subtypes, subtypes)
     |> assign(:selected_subtype, nil)
     |> assign(:changeset, changeset)}
  end

  def handle_event("select_subtype", %{"real_estate" => %{"subtype" => subtype_label}}, socket) do
    changeset = Ecto.Changeset.put_change(socket.assigns.changeset, :subtype, subtype_label)
    {:noreply, assign(socket, changeset: changeset, selected_subtype: subtype_label)}
  end

  def handle_event("validate", %{"real_estate" => params}, socket) do
    changeset =
      RealEstate.changeset(socket.assigns.changeset.data, params)
      |> Map.put(:action, :validate)

    {:noreply, assign(socket, :changeset, changeset)}
  end

  def handle_event("save", %{"real_estate" => params}, socket) do
    params = Map.put(params, "created_by", socket.assigns.current_user.id)
    params = Map.put(params, "updated_by", socket.assigns.current_user.id)

    case Pulap.Estate.create_real_estate(params) do
      {:ok, _real_estate} ->
        {:noreply,
         socket
         |> put_flash(:info, "Real estate created successfully")
         |> push_navigate(to: "/real-estates")}

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset, label: "[DEBUG] Changeset errors on save")
        {:noreply, assign(socket, :changeset, changeset)}
    end
  end

  defp get_options_by_set(set_key) do
    Option
    |> where([o], o.set_id == ^get_set_id(set_key))
    |> Repo.all()
    |> Enum.map(&{&1.label, &1.label})
  end

  defp get_options_by_parent_label(_set_key, parent_label) do
    Option
    |> where([o], o.label == ^parent_label)
    |> Repo.all()
    |> Enum.map(&{&1.label, &1.label})
  end

  defp get_set_id(set_key) do
    Pulap.Set.Set
    |> where([s], s.key == ^set_key)
    |> Repo.one()
    |> Map.get(:id)
  end
end