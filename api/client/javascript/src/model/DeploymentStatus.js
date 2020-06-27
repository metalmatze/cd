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

import ApiClient from '../ApiClient';
import DeploymentStepStatus from './DeploymentStepStatus';

/**
 * The DeploymentStatus model module.
 * @module model/DeploymentStatus
 * @version 0.0.0
 */
class DeploymentStatus {
    /**
     * Constructs a new <code>DeploymentStatus</code>.
     * @alias module:model/DeploymentStatus
     * @param steps {Array.<module:model/DeploymentStepStatus>} 
     */
    constructor(steps) { 
        
        DeploymentStatus.initialize(this, steps);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, steps) { 
        obj['steps'] = steps;
    }

    /**
     * Constructs a <code>DeploymentStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DeploymentStatus} obj Optional instance to populate.
     * @return {module:model/DeploymentStatus} The populated <code>DeploymentStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DeploymentStatus();

            if (data.hasOwnProperty('steps')) {
                obj['steps'] = ApiClient.convertToType(data['steps'], [DeploymentStepStatus]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/DeploymentStepStatus>} steps
 */
DeploymentStatus.prototype['steps'] = undefined;






export default DeploymentStatus;

