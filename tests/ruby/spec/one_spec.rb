

require 'rest-client'
require 'json'
require 'securerandom'




HOST = 'localhost:5351'


def endpoint_installer(data)
  response = RestClient.post HOST, data
  expect(response.code).to eql 200
  response
end


def get_helper(path)
  response = RestClient.get HOST + path
  expect(response.code).to eql 200
  response
end

def random_path
  '/' + SecureRandom.uuid
end


RSpec.describe 'restjester' do


  it "can install a GET endpoint and GET it" do

    resource = random_path
    data = {a:1, b:2}.to_json

    endpoint_installer(path:resource, data:data)
    expect(get_helper(resource)).to eql data

  end


  it "can install a GET endpoint with query params and GET it" do

    resource = random_path + '?a=1&b=2'
    data = {y:1, z:2}.to_json

    endpoint_installer(path:resource, data:data)
    expect(get_helper(resource)).to eql data

  end


  it "can install a GET endpoint with query params and GET it even with different order" do

    resource = random_path + '?b=2&a=1'
    data = {y:1, z:2}.to_json

    endpoint_installer(path:resource, data:data)
    expect(get_helper(resource)).to eql data

  end


  it "can install a POST endpoint and POST to it" do

    endpoint = random_path

    # install POST endpoint
    response = endpoint_installer(path: endpoint, method: 'POST')

    # successfully post to endpoint
    response = RestClient.post HOST + endpoint, {}
    expect(response.code).to eql 200

  end

  it "can GET all endpoints" do

    resource1 = random_path
    resource2 = random_path
    data1 = {a:1, b:2}.to_json
    data2 = {c:3, d:4}.to_json

    endpoint_installer(path:resource1, data:data1)
    endpoint_installer(path:resource2, data:data2)

    # dump endpoint
    response = get_helper('/')

    # check if endpoints are there somewhere, not the most conclusive test
    endpoints = JSON.parse(response)
    expect(endpoints.any? { |ep| ep['Data'] == data1 }).to be true
    expect(endpoints.any? { |ep| ep['Data'] == data2 }).to be true
    expect(endpoints.any? { |ep| ep['Data'] == 'bogus' }).to_not be true

  end


  it "can clear all endpoints" do

    resource1 = random_path
    resource2 = random_path
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
    response = RestClient.get HOST
    endpoints = JSON.parse(response)
    expect(JSON.parse(response)).to eql []

    # lets just check for resources just in case dump isn't working, it wont be conclusive either
    expect {
      get_helper(resource1)
      get_helper(resource2)
    }.to raise_exception RestClient::NotFound

  end


  it "endpoint overwrites previous endpoint" do

    resource = random_path
    data1 = {a:1, b:2}.to_json
    data2 = {a:3, b:4}.to_json

    endpoint_installer(path:resource, data:data1)

    expect(get_helper(resource)).to eql data1

    # overwrite
    endpoint_installer(path:resource, data:data2)

    expect(get_helper(resource)).to eql data2

  end


  it "endpoint overwrites previous endpoint ( with query params )" do

    resource = random_path + '?y=1&z=2'
    data1 = {a:1, b:2}.to_json
    data2 = {a:3, b:4}.to_json

    endpoint_installer(path:resource, data:data1)

    expect(get_helper(resource)).to eql data1

    # overwrite
    endpoint_installer(path:resource, data:data2)

    expect(get_helper(resource)).to eql data2

  end


  it "can install and return 404" do

    path = random_path

    endpoint_installer(path:path, status:404)

    expect {
      get_helper(path)
    }.to raise_exception RestClient::NotFound

  end


  specify "README example works" do

    # install resource on restjester
    RestClient.post 'localhost:5351', { path:'/users/1', data: {'username' => 'kswope'}.to_json }

    # GET resource
    response = RestClient.get 'localhost:5351/users/1' 
    expect( JSON.parse( response.body ) ).to eql( { 'username'=>'kswope' } )

  end


end
