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
import DeploymentStatusUpdate from '../model/DeploymentStatusUpdate';
import SetCurrentDeployment from '../model/SetCurrentDeployment';

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


    /**
     * Set the current Deployment
     * @param {module:model/SetCurrentDeployment} setCurrentDeployment 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/Deployment} and HTTP response
     */
    setCurrentDeploymentWithHttpInfo(setCurrentDeployment) {
      let postBody = setCurrentDeployment;
      // verify the required parameter 'setCurrentDeployment' is set
      if (setCurrentDeployment === undefined || setCurrentDeployment === null) {
        throw new Error("Missing the required parameter 'setCurrentDeployment' when calling setCurrentDeployment");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Deployment;
      return this.apiClient.callApi(
        '/deployments/current', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Set the current Deployment
     * @param {module:model/SetCurrentDeployment} setCurrentDeployment 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/Deployment}
     */
    setCurrentDeployment(setCurrentDeployment) {
      return this.setCurrentDeploymentWithHttpInfo(setCurrentDeployment)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Update parts of the Status of a Deployment
     * @param {Number} id 
     * @param {module:model/DeploymentStatusUpdate} deploymentStatusUpdate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/Deployment} and HTTP response
     */
    updateDeploymentStatusWithHttpInfo(id, deploymentStatusUpdate) {
      let postBody = deploymentStatusUpdate;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling updateDeploymentStatus");
      }
      // verify the required parameter 'deploymentStatusUpdate' is set
      if (deploymentStatusUpdate === undefined || deploymentStatusUpdate === null) {
        throw new Error("Missing the required parameter 'deploymentStatusUpdate' when calling updateDeploymentStatus");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Deployment;
      return this.apiClient.callApi(
        '/deployments/{id}/status', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Update parts of the Status of a Deployment
     * @param {Number} id 
     * @param {module:model/DeploymentStatusUpdate} deploymentStatusUpdate 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/Deployment}
     */
    updateDeploymentStatus(id, deploymentStatusUpdate) {
      return this.updateDeploymentStatusWithHttpInfo(id, deploymentStatusUpdate)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}
