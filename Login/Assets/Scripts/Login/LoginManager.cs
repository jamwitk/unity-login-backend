using System;
using System.Collections;
using System.Collections.Generic;
using Data;
using Helpers;
using TMPro;
using UnityEngine;
using UnityEngine.Networking;

namespace Login
{
    public class LoginManager : MonoBehaviour
    {
        [SerializeField] private LoginData loginData;
        [SerializeField] private TMP_InputField usernameInput;
        [SerializeField] private TMP_InputField passwordInput;
        private void Start()
        {
            Crypto.Encrypt("password", "key");
        }
        public void OnClickLogin()
        {
            string username = usernameInput.text;
            string password = passwordInput.text;
            var encryptedPassword = Helpers.Crypto.Encrypt(password, "key");
            StartCoroutine(SendLoginRequest(username, encryptedPassword));
            
        }

        private IEnumerator SendLoginRequest(string username, string password)
        {
            using var request = UnityWebRequest.Get($"{loginData.url}:{loginData.port}/login?username={username}&password={password}");
            yield return request.SendWebRequest();
            if (request.result is UnityWebRequest.Result.ConnectionError or UnityWebRequest.Result.ProtocolError)
            {
                Debug.LogError("Network Error: " + request.error);
            }
            else
            {
                ProcessLoginResponse(request.downloadHandler.text);
            }
        }
        private void ProcessLoginResponse(string response)
        {
            Debug.Log("repsonse: "+response);
            if (response == "success")
            {
                Debug.Log("Login Successful");
            }
            else
            {
                Debug.Log("Login Failed");
            }
        }
    }
}
