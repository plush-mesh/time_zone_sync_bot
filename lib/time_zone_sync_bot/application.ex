defmodule TimeZoneSyncBot.Application do
  use Application

  @impl true
  def start(_type, _args) do
    webhook_config = [
      host: Application.fetch_env!(:time_zone_sync_bot, :host),
      local_port: Application.fetch_env!(:time_zone_sync_bot, :local_port)
    ]

    bot_config = [
      token: Application.fetch_env!(:time_zone_sync_bot, :token)
    ]

    children = [
      TimeZoneSyncBot.Repo,
      {Tz.UpdatePeriodically, []},
      {Telegram.Webhook, config: webhook_config, bots: [{TimeZoneSyncBot.Bot, bot_config}]}
    ]

    opts = [strategy: :one_for_one, name: TimeZoneSyncBot.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
