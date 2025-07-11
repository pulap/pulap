defmodule PulapWeb.Router do
  use PulapWeb, :router

  import PulapWeb.UserAuth

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, html: {PulapWeb.Layouts, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
    plug :fetch_current_user
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  pipeline :browser_redirect_if_logged_in do
    plug :redirect_if_user_is_authenticated
  end

  pipeline :browser_require_auth do
    plug :require_authenticated_user
  end

  scope "/", PulapWeb do
    pipe_through :browser

    get "/", HomeController, :show
  end

  # Other scopes may use custom stacks.
  # scope "/api", PulapWeb do
  #   pipe_through :api
  # end

  # Enable LiveDashboard and Swoosh mailbox preview in development
  if Application.compile_env(:pulap, :dev_routes) do
    # If you want to use the LiveDashboard in production, you should put
    # it behind authentication and allow only admins to access it.
    # If your application does not have an admins-only section yet,
    # you can use Plug.BasicAuth to set up some basic authentication
    # as long as you are also using SSL (which you should anyway).
    import Phoenix.LiveDashboard.Router

    scope "/dev" do
      pipe_through :browser

      live_dashboard "/dashboard", metrics: PulapWeb.Telemetry
      forward "/mailbox", Plug.Swoosh.MailboxPreview
    end
  end

  ## Authentication routes

  scope "/", PulapWeb do
    pipe_through [:browser, :redirect_if_user_is_authenticated]

    get "/users/register", UserRegistrationController, :new
    post "/users/register", UserRegistrationController, :create
    get "/users/log_in", UserSessionController, :new
    post "/users/log_in", UserSessionController, :create
    get "/users/reset_password", UserResetPasswordController, :new
    post "/users/reset_password", UserResetPasswordController, :create
    get "/users/reset_password/:token", UserResetPasswordController, :edit
    put "/users/reset_password/:token", UserResetPasswordController, :update
  end

  # Place logout routes BEFORE the authenticated user scope to avoid shadowing by resources "/users"
  scope "/", PulapWeb do
    pipe_through [:browser]

    get "/users/log_out", UserSessionController, :not_found
    delete "/users/log_out", UserSessionController, :delete
    get "/users/confirm", UserConfirmationController, :new
    post "/users/confirm", UserConfirmationController, :create
    get "/users/confirm/:token", UserConfirmationController, :edit
    post "/users/confirm/:token", UserConfirmationController, :update
  end

  scope "/", PulapWeb do
    pipe_through [:browser, :require_authenticated_user]

    get "/users/settings", UserSettingsController, :edit
    put "/users/settings", UserSettingsController, :update
    get "/users/settings/confirm_email/:token", UserSettingsController, :confirm_email

    resources "/users", UserController
    resources "/roles", RoleController
    resources "/permissions", PermissionController
    resources "/resources", ResourceController

    resources "/teams", TeamController do
      get "/members", TeamController, :members, as: :members
      post "/assign_member", TeamController, :assign_member, as: :assign_member
      delete "/members/:id", TeamController, :delete_member, as: :delete_member
      get "/members/:id/roles", TeamController, :member_roles, as: :member_roles

      post "/members/:id/assign_role", TeamController, :assign_member_role,
        as: :assign_member_role

      delete "/members/:id/roles/:role_id", TeamController, :revoke_member_role,
        as: :revoke_member_role
    end

    get "/organizations/default", OrganizationController, :show_single

    resources "/organizations", OrganizationController, except: [:new, :create] do
      get "/owners", OrganizationController, :owners
      post "/assign_owner", OrganizationController, :assign_owner
      delete "/owners/:id", OrganizationController, :revoke_owner
    end


    get "/users/:id/roles", UserController, :roles
    get "/users/:id/permissions", UserController, :permissions
    # (to implement)
    post "/users/:id/assign_role", UserController, :assign_role
    # (to implement)
    delete "/users/:id/roles/:role_id", UserController, :revoke_role
    post "/users/:id/assign_permission", UserController, :assign_permission
    delete "/users/:id/permissions/:permission_id", UserController, :revoke_permission

    get "/roles/:id/permissions", RoleController, :permissions
    post "/roles/:id/assign_permission", RoleController, :assign_permission
    delete "/roles/:id/permissions/:permission_id", RoleController, :revoke_permission

    get "/resources/:id/permissions", ResourceController, :permissions
    post "/resources/:id/assign_permission", ResourceController, :assign_permission
    delete "/resources/:id/permissions/:permission_id", ResourceController, :revoke_permission
    
    resources "/sets", SetController do
      resources "/options", OptionController,
        only: [:index, :new, :create, :edit, :update, :delete, :show],
        as: :option
    end
    
    resources "/real-estates", RealEstateController, except: [:new, :show]
    live "/real-estates/new", RealEstateFormLive, :new
    get "/real-estates/:id", RealEstateController, :show
  end
end
