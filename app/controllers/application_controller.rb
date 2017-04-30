class ApplicationController < ActionController::API

  private

  def plaid_client
    PlaidClient.instance.plaid
  end

  def render_error(error, status)
    render json: error, status: status
  end
end
