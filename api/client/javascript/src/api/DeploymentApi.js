/**
 * SignalCD
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import Deployment from '../model/Deployment';

/**
* Deployment service.
* @module api/DeploymentApi
* @version 0.0.0
*/
export default class DeploymentApi {

    /**
    * Constructs a new DeploymentApi. 
    * @alias module:api/DeploymentApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }



    /**
     * Get the current Deployment
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/Deployment} and HTTP response
     */
    getCurrentDeploymentWithHttpInfo() {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Deployment;
      return this.apiClient.callApi(
        '/deployments/current', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get the current Deployment
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/Deployment}
     */
    getCurrentDeployment() {
      return this.getCurrentDeploymentWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * List Deployments
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link Array.<module:model/Deployment>} and HTTP response
     */
    listDeploymentsWithHttpInfo() {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [Deployment];
      return this.apiClient.callApi(
        '/deployments', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * List Deployments
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link Array.<module:model/Deployment>}
     */
    listDeployments() {
      return this.listDeploymentsWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}
