

require 'rest-client'
require 'json'
require 'securerandom'




HOST = 'localhost:5351'


def get_helper(path = '/')
  response = RestClient.get HOST + path
  expect(response.code).to eql 200
  response
end

def endpoint_installer(data)
  response = RestClient.post HOST, data
  expect(response.code).to eql 200
  response
end


RSpec.describe 'restjester' do

  let(:path) { '/findme' }
  let(:host){ 'localhost:5351' }
  let(:data1){ {a:1, b:2, c:3}.to_json }
  let(:data2){ {a:4, b:5, c:6}.to_json }


  it "can install a GET endpoint and GET it" do

    resource = '/myresource/' + SecureRandom.uuid

    # install endpoint
    endpoint_installer(path:resource, data:data1)

    # get endpoint
    response = get_helper(resource)
    expect(response).to eql data1

  end


  it "can install a POST endpoint and POST to it" do

    endpoint = '/myendpoint/' + SecureRandom.uuid

    # install endpoint
    response = RestClient.post host, { path: endpoint, method: 'POST' }
    response = endpoint_installer(path: endpoint, method: 'POST')
    expect(response.code).to eql 200

    # post to endpoint
    response = RestClient.post host + endpoint, {bogus:'data'}
    expect(response.code).to eql 200

  end

  it "can GET all endpoints" do

    resource1 = '/myresource/' + SecureRandom.uuid
    resource2 = '/myresource/' + SecureRandom.uuid
    data1 = {a:1, b:2}.to_json
    data2 = {c:3, d:4}.to_json

    endpoint_installer(path:resource1, data:data1)
    endpoint_installer(path:resource2, data:data2)

    # dump endpoint
    response = get_helper()

    # check if endpoint is there somewhere, not the most conclusive test
    endpoints = JSON.parse(response)
    expect(endpoints.any? { |ep| ep['Data'] == data1 }).to be true
    expect(endpoints.any? { |ep| ep['Data'] == data2 }).to be true
    expect(endpoints.any? { |ep| ep['Data'] == 'bogus' }).to_not be true

  end


  it "can clear all endpoints" do

    resource1 = '/myresource/' + SecureRandom.uuid
    resource2 = '/myresource/' + SecureRandom.uuid
    data1 = {a:5, b:6}.to_json
    data2 = {c:7, d:8}.to_json

    endpoint_installer(path:resource1, data:data1)
    endpoint_installer(path:resource2, data:data2)

    # make sure they are there
    expect {
      get_helper(resource1)
      get_helper(resource2)
    }.to_not raise_exception

    # now delete them
    response = RestClient.delete HOST
    expect(response.code).to eql 200

    # how else can we test clear other than rely on GET all working correctly??
    response = get_helper()
    endpoints = JSON.parse(response)
    expect(JSON.parse(response)).to eql []

    # lets just check for data1 just in case dump isn't working, it wont be conclusive however
    expect {
      get_helper(resource1)
      get_helper(resource2)
    }.to raise_exception RestClient::NotFound

  end


  it "endpoint overwrites previous endpoint" do

    response = RestClient.post host, { path: path, data: data1 }
    expect(response.code).to eql 200

    # overwrite
    response = RestClient.post host, { path: path, data: data2 }
    expect(response.code).to eql 200

    # GET endpoint
    response = RestClient.get host + path
    expect(response.code).to eql 200
    expect(response).to eql data2

  end


  it "can install and return 404" do

    response = RestClient.post host, { path: path, status: 404 }
    expect(response.code).to eql 200

    expect {
      RestClient.get host + path
    }.to raise_exception RestClient::NotFound

  end


  specify "README example works" do

    # install resource on restjester
    RestClient.post 'localhost:5351', { path:'/users/1', data: {username: 'kswope'}.to_json }

    # GET resource
    response = RestClient.get 'localhost:5351/users/1' 
    expect( JSON.parse( response.body ) ).to eql( { 'username'=>'kswope' } )

  end


  specify "doc example works" do

    RestClient.post 'localhost:5351', { path:'/users/1', data: {username: 'kswope'}.to_json }
    response = RestClient.get 'localhost:5351/users/1' 
    expect( JSON.parse( response.body ) ).to eql( { 'username'=>'kswope' } )

    RestClient.post 'localhost:5351', { method:'PUT', path:'/users/1' }
    response = RestClient.get 'localhost:5351/users/1' 
    expect( response.code ).to eql 200

    RestClient.post 'localhost:5351', { method:'DELETE', path:'/users/1', status:403 }
    expect {
      response = RestClient.delete 'localhost:5351/users/1'
    }.to raise_exception RestClient::Forbidden

  end


end
