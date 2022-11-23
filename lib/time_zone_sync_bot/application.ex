defmodule TimeZoneSyncBot.Application do
  use Application

  @impl true
  def start(_type, _args) do
    children = [
      TimeZoneSyncBot.Repo,
      {Tz.UpdatePeriodically, []}
    ]

    opts = [strategy: :one_for_one, name: TimeZoneSyncBot.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
