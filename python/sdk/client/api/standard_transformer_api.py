# coding: utf-8

"""
    Merlin

    API Guide for accessing Merlin's model management, deployment, and serving functionalities  # noqa: E501

    OpenAPI spec version: 0.14.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import re  # noqa: F401

# python 2 and python 3 compatibility library
import six

from client.api_client import ApiClient


class StandardTransformerApi(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    Ref: https://github.com/swagger-api/swagger-codegen
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

    def standard_transformer_simulate_post(self, **kwargs):  # noqa: E501
        """Simulate standard transformer  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.standard_transformer_simulate_post(async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param StandardTransformerSimulationRequest body:
        :return: VersionEndpoint
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.standard_transformer_simulate_post_with_http_info(**kwargs)  # noqa: E501
        else:
            (data) = self.standard_transformer_simulate_post_with_http_info(**kwargs)  # noqa: E501
            return data

    def standard_transformer_simulate_post_with_http_info(self, **kwargs):  # noqa: E501
        """Simulate standard transformer  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.standard_transformer_simulate_post_with_http_info(async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param StandardTransformerSimulationRequest body:
        :return: VersionEndpoint
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['body']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method standard_transformer_simulate_post" % key
                )
            params[key] = val
        del params['kwargs']

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'body' in params:
            body_params = params['body']
        # Authentication setting
        auth_settings = ['Bearer']  # noqa: E501

        return self.api_client.call_api(
            '/standard_transformer/simulate', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='VersionEndpoint',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)
